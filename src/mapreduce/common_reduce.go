package mapreduce

import (
	"io/ioutil"
	"encoding/json"
	"sort"
)

type keyListOfValues struct {
	Key string
	Values []string
}

// doReduce does the job of a reduce worker: it reads the intermediate
// key/value pairs (produced by the map phase) for this task, sorts the
// intermediate key/value pairs by key, calls the user-defined reduce function
// (reduceF) for each key, and writes the output to disk.
func doReduce(
	jobName string, // the name of the whole MapReduce job
	reduceTaskNumber int, // which reduce task this is
	nMap int, // the number of map tasks that were run ("M" in the paper)
	reduceF func(key string, values []string) string,
) {
	var resultKeyValue []KeyValue
	for i:= 0; i<nMap; i++ {
		fileName := reduceName(jobName,i,reduceTaskNumber)
		contents,err := ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		contentsString := string(contents)
		var jsonContent []KeyValue
		err = json.Unmarshal([]byte(contentsString), &jsonContent)
		if err != nil {
			panic(err)
		}
		resultKeyValue = append(resultKeyValue,jsonContent...)
	}
	sort.Slice(resultKeyValue, func(i,j int) bool {
		return resultKeyValue[i].Key < resultKeyValue[j].Key
	})

	foldOfKeys := make([]keyListOfValues,0)
	for k := 0; k <len(resultKeyValue); k++ {
		key := resultKeyValue[k].Key
		value := resultKeyValue[k].Value
		if len(foldOfKeys) == 0 {
			newValue := keyListOfValues{key,[]string{value}}
			foldOfKeys = append(foldOfKeys,newValue)
		} else {
			lastElement := foldOfKeys[len(foldOfKeys) -1]
			if key == lastElement.Key {
				foldOfKeys[len(foldOfKeys)-1].Values = append(lastElement.Values, value)
			} else {
				foldOfKeys = append(foldOfKeys, keyListOfValues{key,[]string{value}})
			}
		}
	}
	resultStrings := make([]string,0)
	//apply reducef and write on file(how to combine the output?)
	for index2 := 0; index2<len(foldOfKeys); index2++ {
		ind := foldOfKeys[index2].Key
		val := foldOfKeys[index2].Values
		newString := reduceF(ind,val)
		resultStrings=append(resultStrings,newString)
	}

	// TODO:
	// You will need to write this function.
	// You can find the intermediate file for this reduce task from map task number
	// m using reduceName(jobName, m, reduceTaskNumber).
	// Remember that you've encoded the values in the intermediate files, so you
	// will need to decode them. If you chose to use JSON, you can read out
	// multiple decoded values by creating a decoder, and then repeatedly calling
	// .Decode() on it until Decode() returns an error.
	//
	// You should write the reduced output in as JSON encoded KeyValue
	// objects to a file named mergeName(jobName, reduceTaskNumber). We require
	// you to use JSON here because that is what the merger than combines the
	// output from all the reduce tasks expects. There is nothing "special" about
	// JSON -- it is just the marshalling format we chose to use. It will look
	// something like this:
	//
	// enc := json.NewEncoder(mergeFile)
	// for key in ... {
	// 	enc.Encode(KeyValue{key, reduceF(...)})
	// }
	// file.Close()
}
