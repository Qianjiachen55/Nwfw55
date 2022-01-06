package env

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/Qianjiachen55/Nwfw55/framework/contract"
	"io"
	"os"
	"path"
	"strings"
)

type NwfwEnv struct {
	folder string
	maps map[string]string
}


func NewNwfwEnv(params ...interface{}) (interface{},error) {
	if len(params) != 1{
		return nil,errors.New("NewNwfw param error")
	}

	folder := params[0].(string)

	nwfwEnv := &NwfwEnv{
		folder: folder,
		maps : map[string]string{"APP_ENV":contract.EnvDevelopment},
	}

	file := path.Join(folder,".env")
	_,err := os.Stat(file)
	if err == nil{
		fi,err := os.Open(file)
		if err ==nil{
			defer fi.Close()

			br := bufio.NewReader(fi)

			for {
				line ,_,c := br.ReadLine()
				if c == io.EOF{
					break
				}
				s := bytes.SplitN(line,[]byte{'='},2)
				if len(s)<2{
					continue
				}
				key :=string(s[0])
				val := string(s[1])
				nwfwEnv.maps[key] = val
			}
		}
	}

	for _,e := range os.Environ(){
		pair := strings.SplitN(e,"=",2)
		if len(pair)<2{
			continue
		}
		nwfwEnv.maps[pair[0]] = pair[1]
	}

	return nwfwEnv, nil

}

func (env *NwfwEnv) AppEnv() string{
	return env.Get("APP_ENV")
}

func (env *NwfwEnv) IsExist(key string) bool{
	_,ok := env.maps[key]
	return ok
}

func (env *NwfwEnv) Get(key string) string{
	if val,ok :=env.maps[key];ok{
		return val
	}
	return ""
}

func (env *NwfwEnv) All() map[string]string{
	return env.maps
}
