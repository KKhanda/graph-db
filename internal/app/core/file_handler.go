package core

import (
	"fmt"
	"os"
	"path/filepath"
	"errors"
	"graph-db/internal/pkg/utils"
	"io/ioutil"
	"strings"
	"strconv"
	"graph-db/internal/app/core/globals"
)

var (
	rootPath = "databases"
	err error
)

type FileHandler struct {
}

func (fh FileHandler) InitFileSystem() {
	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		os.Mkdir(rootPath, os.ModePerm)
	}
}

func (fh FileHandler) InitDatabaseStructure(dbTitle string) {
	var dbPath = filepath.Join(rootPath, dbTitle)
	var storagePath = filepath.Join(rootPath, dbTitle, "storage")
	var nodesPath = filepath.Join(storagePath, "nodes")
	var nodesIdPath = filepath.Join(nodesPath, "id")
	var nodesStorePath = filepath.Join(nodesPath, "store")
	var relationshipsPath = filepath.Join(storagePath, "relationships")
	var relationshipsIdPath = filepath.Join(relationshipsPath, "id")
	var relationshipsStorePath = filepath.Join(relationshipsPath, "store")
	var propertiesPath = filepath.Join(storagePath, "properties")
	var propertiesIdPath = filepath.Join(propertiesPath, "id")
	var propertiesStorePath = filepath.Join(propertiesPath, "store")

	os.MkdirAll(nodesIdPath, os.ModePerm)
	os.MkdirAll(nodesStorePath, os.ModePerm)
	os.MkdirAll(relationshipsIdPath, os.ModePerm)
	os.MkdirAll(relationshipsStorePath, os.ModePerm)
	os.MkdirAll(propertiesIdPath, os.ModePerm)
	os.MkdirAll(propertiesStorePath, os.ModePerm)

	// nodes/id
	globals.NodesId, err = os.Create(filepath.Join(nodesIdPath, "nodes.id"))
	utils.CheckError(err)
	globals.LabelsId, err = os.Create(filepath.Join(nodesIdPath, "labels.id"))
	utils.CheckError(err)
	globals.LabelsTitlesId, err = os.Create(filepath.Join(nodesIdPath, "labelsTitles.id"))
	utils.CheckError(err)
	// nodes/store
	globals.NodesStore, err = os.Create(filepath.Join(nodesStorePath, "nodes.store"))
	utils.CheckError(err)
	globals.LabelsStore, err = os.Create(filepath.Join(nodesStorePath, "labels.store"))
	utils.CheckError(err)
	globals.LabelsTitlesStore, err = os.Create(filepath.Join(nodesStorePath, "labelsTitles.store"))
	utils.CheckError(err)

	// relationships/id
	globals.RelationshipsId, err = os.Create(filepath.Join(relationshipsIdPath, "relationships.id"))
	utils.CheckError(err)
	globals.RelationshipsTitlesId, err = os.Create(filepath.Join(relationshipsIdPath, "relationshipsTitles.id"))
	utils.CheckError(err)
	// relationships/store
	globals.RelationshipsStore, err = os.Create(filepath.Join(relationshipsStorePath, "relationships.store"))
	utils.CheckError(err)
	globals.RelationshipsTitlesStore, err = os.Create(filepath.Join(relationshipsStorePath, "relationshipsTitles.store"))
	utils.CheckError(err)

	// properties/id
	globals.PropertiesId, err = os.Create(filepath.Join(propertiesIdPath, "properties.id"))
	utils.CheckError(err)
	globals.PropertiesTitlesId, err = os.Create(filepath.Join(propertiesIdPath, "propertiesTitles.id"))
	utils.CheckError(err)
	globals.StringId, err = os.Create(filepath.Join(propertiesIdPath, "string.id"))
	utils.CheckError(err)
	globals.DoubleId, err = os.Create(filepath.Join(propertiesIdPath, "double.id"))
	utils.CheckError(err)
	// properties/store
	globals.PropertiesStore, err = os.Create(filepath.Join(propertiesStorePath, "properties.store"))
	utils.CheckError(err)
	globals.PropertiesTitlesStore, err = os.Create(filepath.Join(propertiesStorePath, "propertiesTitles.store"))
	utils.CheckError(err)
	globals.StringStore, err = os.Create(filepath.Join(propertiesStorePath, "string.store"))
	utils.CheckError(err)
	globals.DoubleStore, err = os.Create(filepath.Join(propertiesStorePath, "double.store"))
	utils.CheckError(err)

	//config
	globals.Config, err = os.Create(filepath.Join(dbPath, "connections.config"))
	utils.CheckError(err)

	globals.NodesId.WriteString(fmt.Sprintf("%d", 0))
	globals.LabelsId.WriteString(fmt.Sprintf("%d", 0))
	globals.LabelsTitlesId.WriteString(fmt.Sprintf("%d", 0))

	globals.RelationshipsId.WriteString(fmt.Sprintf("%d", 0))
	globals.RelationshipsTitlesId.WriteString(fmt.Sprintf("%d", 0))

	globals.PropertiesId.WriteString(fmt.Sprintf("%d", 0))
	globals.PropertiesTitlesId.WriteString(fmt.Sprintf("%d", 0))
	globals.StringId.WriteString(fmt.Sprintf("%d", 0))
	globals.DoubleId.WriteString(fmt.Sprintf("%d", 0))
}

func (fh FileHandler) SwitchDatabaseStructure(dbTitle string) (err error) {
	if _, err := os.Stat(filepath.Join(rootPath, dbTitle)); err == nil {
		var storagePath = filepath.Join(rootPath, dbTitle, "storage")
		var nodesPath = filepath.Join(storagePath, "nodes")
		var nodesIdPath = filepath.Join(nodesPath, "id")
		var nodesStorePath = filepath.Join(nodesPath, "store")
		var relationshipsPath = filepath.Join(storagePath, "relationships")
		var relationshipsIdPath = filepath.Join(relationshipsPath, "id")
		var relationshipsStorePath = filepath.Join(relationshipsPath, "store")
		var propertiesPath = filepath.Join(storagePath, "properties")
		var propertiesIdPath = filepath.Join(propertiesPath, "id")
		var propertiesStorePath = filepath.Join(propertiesPath, "store")

		// nodes/id
		globals.NodesId, err = os.OpenFile(filepath.Join(nodesIdPath, "nodes.id"), os.O_RDWR, 0666)
		utils.CheckError(err)
		globals.LabelsId, err = os.OpenFile(filepath.Join(nodesIdPath, "labels.id"), os.O_RDWR, 0666)
		utils.CheckError(err)
		globals.LabelsTitlesId, err = os.OpenFile(filepath.Join(nodesIdPath, "labelsTitles.id"), os.O_RDWR, 0666)
		utils.CheckError(err)
		// nodes/store
		globals.NodesStore, err = os.OpenFile(filepath.Join(nodesStorePath, "nodes.store") , os.O_RDWR, 0666)
		utils.CheckError(err)
		globals.LabelsStore, err = os.OpenFile(filepath.Join(nodesStorePath, "labels.store"), os.O_RDWR, 0666)
		utils.CheckError(err)
		globals.LabelsTitlesStore, err = os.OpenFile(filepath.Join(nodesStorePath, "labelsTitles.store"), os.O_RDWR, 0666)
		utils.CheckError(err)

		// relationships/id
		globals.RelationshipsId, err = os.OpenFile(filepath.Join(relationshipsIdPath, "relationships.id"), os.O_RDWR, 0666)
		utils.CheckError(err)
		globals.RelationshipsTitlesId, err = os.OpenFile(filepath.Join(relationshipsIdPath, "relationshipsTitles.id"), os.O_RDWR, 0666)
		utils.CheckError(err)
		// relationships/store
		globals.RelationshipsStore, err = os.OpenFile(filepath.Join(relationshipsStorePath, "relationships.store"), os.O_RDWR, 0666)
		utils.CheckError(err)
		globals.RelationshipsTitlesStore, err = os.OpenFile(filepath.Join(relationshipsStorePath, "relationshipsTitles.store"), os.O_RDWR, 0666)
		utils.CheckError(err)

		// properties/id
		globals.PropertiesId, err = os.OpenFile(filepath.Join(propertiesIdPath, "properties.id"), os.O_RDWR, 0666)
		utils.CheckError(err)
		globals.PropertiesTitlesId, err = os.OpenFile(filepath.Join(propertiesIdPath, "propertiesTitles.id"), os.O_RDWR, 0666)
		utils.CheckError(err)
		globals.StringId, err = os.OpenFile(filepath.Join(propertiesIdPath, "string.id"), os.O_RDWR, 0666)
		utils.CheckError(err)
		globals.DoubleId, err = os.OpenFile(filepath.Join(propertiesIdPath, "double.id"), os.O_RDWR, 0666)
		utils.CheckError(err)
		// properties/store
		globals.PropertiesStore, err = os.OpenFile(filepath.Join(propertiesStorePath, "properties.store"), os.O_RDWR, 0666)
		utils.CheckError(err)
		globals.PropertiesTitlesStore, err = os.OpenFile(filepath.Join(propertiesStorePath, "propertiesTitles.store"), os.O_RDWR, 0666)
		utils.CheckError(err)
		globals.StringStore, err = os.OpenFile(filepath.Join(propertiesStorePath, "string.store"), os.O_RDWR, 0666)
		utils.CheckError(err)
		globals.DoubleStore, err = os.OpenFile(filepath.Join(propertiesStorePath, "double.store"), os.O_RDWR, 0666)
		utils.CheckError(err)

		globals.Config, err = os.OpenFile(filepath.Join(rootPath, dbTitle, "connections.config"), os.O_RDWR, 0666)

		return err
	} else {
		return errors.New(fmt.Sprintf("Database with title %s does not exist", dbTitle))
	}
}

func (fh FileHandler) DropDatabase(dbTitle string) (err error) {
	if _, err := os.Stat(filepath.Join(rootPath, dbTitle)); err == nil {
		if globals.CurrentDb == dbTitle {
			// nodes/id
			err = globals.NodesId.Close()
			globals.NodesId = nil
			utils.CheckError(err)
			err = globals.LabelsId.Close()
			globals.LabelsId = nil
			utils.CheckError(err)
			err = globals.LabelsTitlesId.Close()
			globals.LabelsTitlesId = nil
			utils.CheckError(err)
			// nodes/store
			err = globals.NodesStore.Close()
			globals.NodesStore = nil
			utils.CheckError(err)
			err = globals.LabelsStore.Close()
			globals.LabelsStore = nil
			utils.CheckError(err)
			err = globals.LabelsTitlesStore.Close()
			globals.LabelsTitlesStore = nil
			utils.CheckError(err)

			// relationships/id
			err = globals.RelationshipsId.Close()
			globals.RelationshipsId = nil
			utils.CheckError(err)
			err = globals.RelationshipsTitlesId.Close()
			globals.RelationshipsId = nil
			utils.CheckError(err)
			// relationships/store
			err = globals.RelationshipsStore.Close()
			globals.RelationshipsStore = nil
			utils.CheckError(err)
			err = globals.RelationshipsTitlesStore.Close()
			globals.RelationshipsTitlesStore = nil
			utils.CheckError(err)

			// properties/id
			err = globals.PropertiesId.Close()
			globals.PropertiesId = nil
			utils.CheckError(err)
			err = globals.PropertiesTitlesId.Close()
			globals.PropertiesTitlesId = nil
			utils.CheckError(err)
			err = globals.StringId.Close()
			globals.StringId = nil
			utils.CheckError(err)
			err = globals.DoubleId.Close()
			globals.DoubleId = nil
			utils.CheckError(err)
			// properties/store
			err = globals.PropertiesStore.Close()
			globals.PropertiesStore = nil
			utils.CheckError(err)
			err = globals.PropertiesTitlesStore.Close()
			globals.PropertiesTitlesStore = nil
			utils.CheckError(err)
			err = globals.StringStore.Close()
			globals.StringStore = nil
			utils.CheckError(err)
			err = globals.DoubleStore.Close()
			globals.DoubleStore = nil
			utils.CheckError(err)
			// config
			err = globals.Config.Close()
			globals.Config = nil
			utils.CheckError(err)
		}
		err = os.RemoveAll(filepath.Join(rootPath, dbTitle))
		return err
	} else {
		return errors.New(fmt.Sprintf("Database with title %s does not exist", dbTitle))
	}
}

func (fh FileHandler) Read(file *os.File, offset int, bs *[]byte, id int) (err error) {
	//if strings.HasSuffix(file.Name(), "nodes.store") {
	//	file = globals.NodesStore
	//}
	//if strings.HasSuffix(file.Name(), "labels.store") {
	//	file = globals.LabelsStore
	//}
	//if strings.HasSuffix(file.Name(), "relationships.store") {
	//	file = globals.RelationshipsStore
	//}
	//if strings.HasSuffix(file.Name(), "properties.store") {
	//	file = globals.PropertiesStore
	//}
	//if strings.HasSuffix(file.Name(), "string.store") {
	//	file = globals.StringStore
	//}
	//if strings.HasSuffix(file.Name(), "double.store") {
	//	file = globals.DoubleStore
	//}
	bytesRead, err := file.ReadAt(*bs, int64(offset))
	if bytesRead != len(*bs) {
		err = errors.New("read: read less bytes than expected")
	}
	return err
}
func (fh FileHandler) Write(file *os.File, offset int, bs []byte, id int) (err error) {
	//if strings.HasSuffix(file.Name(), "nodes.store") {
	//	file = globals.NodesStore
	//}
	//if strings.HasSuffix(file.Name(), "labels.store") {
	//	file = globals.LabelsStore
	//}
	//if strings.HasSuffix(file.Name(), "relationships.store") {
	//	file = globals.RelationshipsStore
	//}
	//if strings.HasSuffix(file.Name(), "properties.store") {
	//	file = globals.PropertiesStore
	//}
	//if strings.HasSuffix(file.Name(), "string.store") {
	//	file = globals.StringStore
	//}
	//if strings.HasSuffix(file.Name(), "double.store") {
	//	file = globals.DoubleStore
	//}
	bytesWritten, err := file.WriteAt(bs, int64(offset))
	if bytesWritten != len(bs) {
		err = errors.New("write: wrote less bytes than expected")
	}
	return err
}

func (fh FileHandler) ReadId(file *os.File) (id int, err error) {
	fileData, err := ioutil.ReadFile(file.Name())
	if err == nil {
		ids := strings.Split(string(fileData), "\n")
		id, err := strconv.Atoi(ids[0])
		if err == nil {
			if len(ids) == 1 {
				str := strconv.Itoa(id + 1)
				err := ioutil.WriteFile(file.Name(), []byte(str), os.ModePerm)
				if err == nil {
					return id, err
				}
			} else {
				str := strings.Join(ids[1:], "\n")
				err := ioutil.WriteFile(file.Name(), []byte(str), os.ModePerm)
				if err == nil {
					return id, err
				}
			}
		}
	}
	return 0, err
}

func (fh FileHandler) FreeId(file *os.File, id int) (err error) {
	fileData, err := ioutil.ReadFile(file.Name())
	if err == nil {
		ids := strings.Split(string(fileData), "\n")
		str := strconv.Itoa(id)
		var fetchedPrev, fetchedNext, firstId, lastId int
		firstId, err = strconv.Atoi(ids[0])
		utils.CheckError(err)
		lastId, err = strconv.Atoi(ids[len(ids) - 1])
		utils.CheckError(err)
		if id < firstId {
			ids = append([]string{str}, ids[:]...)
		} else if id > lastId {
			return errors.New("Bad id (specified id is out of range)")
		} else if id == lastId {
			return errors.New("Bad id (specified id is already free)")
		} else {
			for i := 0; i < len(ids) - 1; i++ {
				fetchedPrev, err = strconv.Atoi(ids[i])
				utils.CheckError(err)
				fetchedNext, err = strconv.Atoi(ids[i + 1])
				utils.CheckError(err)
				if id == fetchedPrev {
					return errors.New("Bad id (specified id is already free)")
				}
				if id > fetchedPrev && id < fetchedNext {
					ids = append(ids[:i + 1], append([]string{str}, ids[i + 1:]...)...)
					break
				}
			}
		}
		str = strings.Join(ids, "\n")
		err = ioutil.WriteFile(file.Name(), []byte(str), os.ModePerm)
	}
	return err
}