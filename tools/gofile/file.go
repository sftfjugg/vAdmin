// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gofile

import (
    "bufio"
    "io"
    "io/ioutil"
    "os"
    "strings"
)

func CreateFile(filePath string, data []byte, perm os.FileMode) error {
    return ioutil.WriteFile(filePath, data, perm)
}

func ReadlineAddHead(filePath string, headstr string) error {
    var result string
    f, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer f.Close()
    rd := bufio.NewReader(f)
    for {
        line, err := rd.ReadString('\n')
        line = strings.Replace(line,"\r","",-1)
        line = strings.Replace(line,"\n","",-1)
        //fmt.Println(line)
        result += headstr + " " + string(line) + "\n"
        if err != nil || io.EOF == err {
            break
        }
    }

    fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)//os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
    w := bufio.NewWriter(fw)
    w.WriteString(result)
    if err != nil {
        return err
    }
    w.Flush()
    return err
}

func ReadlineAddafter(filePath string, afterstr string) error {
    var result string
    f, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer f.Close()
    rd := bufio.NewReader(f)
    for {
        line, err := rd.ReadString('\n')
        line = strings.Replace(line,"\r","",-1)
        line = strings.Replace(line,"\n","",-1)
        //fmt.Println(line)
        result += string(line) + " " + afterstr + "\n"
        if err != nil || io.EOF == err {
            break
        }
    }

    fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)//os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
    w := bufio.NewWriter(fw)
    w.WriteString(result)
    if err != nil {
        return err
    }
    w.Flush()
    return err
}