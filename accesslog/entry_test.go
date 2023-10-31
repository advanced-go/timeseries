package accesslog

type accessLogV2 struct {
	Duration string
}

func _ExampleScanColumnsTemplate() {
	//log := scanColumnsTemplate[AccessLog](nil)

	//fmt.Printf("test: scanColumnsTemplate[AccessLog](nil) -> %v\n", log)

	//Output:
	//fail
}

func _ExampleScannerInterface_V1() {

	//log, status := scanRowsTemplateV1[AccessLog, AccessLog](nil)
	//fmt.Printf("test: scanRowsTemplateV1() -> [status:%v] [elem:%v] [log:%v] \n", status, reflect.TypeOf(log).Elem(), log[0].DurationString)

	//Output:
	//test: scanRowsTemplateV1() -> [status:OK] [elem:timeseries.AccessLog] [log:SCAN() TEST DURATION STRING]

}

func _ExampleScannerInterface() {
	//log, status := scanRowsTemplate[accessLogV2](nil)

	//log, status := scanRowsTemplate[AccessLog](nil)
	//fmt.Printf("test: scanRowsTemplate() -> [status:%v] [elem:%v] [log:%v] \n", status, reflect.TypeOf(log).Elem(), log[0].DurationString)

	//Output:
	//test: scanRowsTemplateV1() -> [status:OK] [elem:timeseries.AccessLog] [log:SCAN() TEST DURATION STRING]

}
