package main

import "os"
import "strings"
import "fmt"
import . "ufp-env-file/model"

func getEnvs(config Config) {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")


		// SPLIT BY SEPARATOR


		fmt.Println(pair[0] + "==" + pair[1])
	}
}

func main() {
	config := GetConfig()
	os.Setenv("UFP_SIMPLE_TEST", "1")
	os.Setenv("UFP_SIMPLE_TEST1", "2")
	os.Setenv("UFP_SIMPLE_TEST2", "3")
	os.Setenv("UFP_SIMPLE_TEST_X", "4")
	os.Setenv("UFP_SIMPLE_TEST_Y", "5")
	os.Setenv("UFP_SIMPLE_TEST_Z", "6")

	/**
	expected result"

	{

	simple:{
	test:{
	x:"4",
	y:"5",
	z:"6" ,
	},
	test1:"2",
	test2:"3"
	}
	}


	 */
	getEnvs(config);
}
