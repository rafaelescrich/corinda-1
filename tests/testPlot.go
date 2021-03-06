package main

import (
	"github.com/bernardoaraujor/corinda/train"
	"github.com/bernardoaraujor/corinda/plot"
	"runtime"
	"fmt"
	"os"
	"encoding/gob"
	"github.com/bernardoaraujor/corinda/elementary"
	"github.com/bernardoaraujor/corinda/composite"
)

func main(){
	//train.Train("test", 10)

	// -----------------------------------------------------------------------------------------------------------------
	// empty tm
	tm := train.Maps{make(map[string]*elementary.Model), make(map[string]*composite.Model)}

	var tm2 = new(train.Maps)
	err := load("maps/testTrainedMaps.gob", &tm2)
	check(err)
	tm.Merge(tm2)

	/*
	for _, em := range tm.ElementaryMap{
		plot.EMLogLog(*em)
	}
	*/

	plot.EMLogLog(*tm.ElementaryMap["Exact Match:500-worst-passwords"])

	// -----------------------------------------------------------------------------------------------------------------


}

/*
// checks for error
func check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}

// Decode Gob file
func load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

*/