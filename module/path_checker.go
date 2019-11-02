package module

// TODO: check dependency loop
func CheckDeps(path Path, rootModuleName string) Path {
	rootModule := path.findModule(rootModuleName)
	if rootModule == nil {
		panic("unknown module: " + rootModule.GetName() + "@" + rootModule.GetVersion())
	}

	checkList := []Module{rootModule}
	checked := map[string]Module{}

	for len(checkList) > 0 {
		m := checkList[0]
		checkList = checkList[1:]

		if checked[m.GetName()] == nil {
			for _, require := range m.GetInfo().Requires {
				if checked[require.Name] == nil {
					if found := path.findModule(require.Name); found != nil {
						checkList = append(checkList, found)
					} else {
						panic("unknown module: " + require.Name + "@" + require.Version)
					}
				}
			}
			checked[m.GetName()] = m
		}
	}

	checkedList := make([]Module, 0, len(checked))
	for _, m := range checked {
		checkedList = append(checkedList, m)
	}

	checkedPath := Path(checkedList)
	checkedPath.Sort()
	return checkedPath
}
