package GameEngine

import (
	"log"
	"os"
	"encoding/xml"
	Collada14 "github.com/metaleap/go-xsd-pkg/khronos.org/files/collada_schema_1_4_go"
)
type ColladaDoc struct {
	XMLName	xml.Name	`xml:"COLLADA"`
	Collada14.TxsdCollada
}

func ImportColladaFile(filename string) (c *ColladaDoc) {
    file, err := os.Open(filename)
    if err!= nil{
        log.Println(err.Error())   
    }
    
    fi, err := file.Stat()
    
    if err!= nil{
        log.Println(err.Error())   
    }
    
    data := make([]byte, fi.Size())
	count, err := file.Read(data)
	if err != nil {
		log.Println(err.Error())   
	}
	log.Printf("read %d bytes\n", count)
    
    err = xml.Unmarshal(data, &c)
    if err!= nil{
        log.Println(err.Error())     
    }
    
    return
}