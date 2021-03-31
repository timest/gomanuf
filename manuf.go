package manuf

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const hexDigit = "0123456789ABCDEF"

var d map[int]interface{}

const ManufSource = "https://code.wireshark.org/review/gitweb?p=wireshark.git;a=blob_plain;f=manuf"

func init() {
	if _, err := os.Stat(".manuf"); os.IsNotExist(err) {
		if DownloadManuf() != nil {
			panic(err)
		}
	}
	d = make(map[int]interface{})
	err := readLine(".manuf", func(s string) {
		l := strings.Split(s, "\t")
		if len(l) > 2 {
			parse(l[0], l[2])
		}
	})
	if err != nil {
		panic(err)
	}
}

func DownloadManuf() error {
	resp, err := http.Get(ManufSource)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(".manuf")
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func parse(mac, comment string) {
    g := strings.Split(mac, "/")
    m := strings.Split(g[0], ":")
    var b int
    if len(g) != 2 {
        b = 48 - len(m) * 8
    } else {
        b, _ = strconv.Atoi(g[1])
    }
    if _, ok := d[b]; !ok {
        d[b] = make(map[uint64]string)
        
    }
    d[b].(map[uint64]string)[b2uint64(m)] = comment
}
func b2uint64(sList []string) uint64 {
    var t uint64
    for i, b := range sList {
        l := strings.Index(hexDigit, string(b[0]))
        r := strings.Index(hexDigit, string(b[1]))
        t += uint64((l << 4) + r) << uint8((6 - i - 1) * 8)
    }
    
    return t
}

func Search(mac string) string {
    s := strings.Split(strings.ToUpper(mac) , ":")
    bint := b2uint64(s)
    for b := range d {
        k := 48 - b
        bint = (bint >> uint8(k)) << uint8(k)
        if _, ok := d[b].(map[uint64]string)[bint]; ok {
            return d[b].(map[uint64]string)[bint]
        }
    }
    return ""
}

func readLine(fileName string, handler func(string)) error {
    f, err := os.Open(fileName)
    defer f.Close()
    if err != nil {
        return err
    }
    buf := bufio.NewReader(f)
    for {
        line, err := buf.ReadString('\n')
        line = strings.TrimSpace(line)
        handler(line)
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
    }
    return nil
}
