package main

import (
	"graph-db/internal/app/core"
	"graph-db/internal/pkg/utils"
	"graph-db/internal/app/core/globals"
	"log"
)

func main() {
	//err := core.InitDb("asd", "local")
	//err = core.SwitchDb("asd")
	//utils.CheckError(err)

	dbTitle := "asd"
	var dfh core.DistributedFileHandler
	dfh.InitFileSystem()
	err := core.InitDb(dbTitle, "distributed")
	dfh.InitDatabaseStructure(dbTitle)
	if err != nil {
		log.Fatal("Error in initialization of database")
	}

	bs := utils.StringToByteArray("Test")
	dfh.Write(globals.NodesStore, 20, bs, 0)
	newBs := make([]byte, 4)
	dfh.Read(globals.NodesStore, 20, &newBs, 0)

	if string(newBs) != string(bs) {
		log.Fatal("Byte arrays are not the same!")
	} else {
		println("Congratulations!")
	}
}