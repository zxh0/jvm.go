package classfile

/*
Module_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 module_name_index;
    u2 module_flags;
    u2 module_version_index;
    u2 requires_count;
    {   u2 requires_index;
        u2 requires_flags;
        u2 requires_version_index;
    } requires[requires_count];
    u2 exports_count;
    {   u2 exports_index;
        u2 exports_flags;
        u2 exports_to_count;
        u2 exports_to_index[exports_to_count];
    } exports[exports_count];
    u2 opens_count;
    {   u2 opens_index;
        u2 opens_flags;
        u2 opens_to_count;
        u2 opens_to_index[opens_to_count];
    } opens[opens_count];
    u2 uses_count;
    u2 uses_index[uses_count];
    u2 provides_count;
    {   u2 provides_index;
        u2 provides_with_count;
        u2 provides_with_index[provides_with_count];
    } provides[provides_count];
}
*/

type ModuleAttribute struct {
	ModuleNameIndex    uint16
	ModuleFlags        uint16
	ModuleVersionIndex uint16
	RequiresTable      []ModuleRequires
	ExportsTable       []ModuleExports
	OpensTable         []ModuleOpens
	UsesIndexTable     []uint16
	ProvidesTable      []ModuleProvides
}

type ModuleRequires struct {
	RequiresIndex        uint16
	RequiresFlags        uint16
	RequiresVersionIndex uint16
}

type ModuleExports struct {
	ExportsIndex        uint16
	ExportsFlags        uint16
	ExportsToIndexTable []uint16
}

type ModuleOpens struct {
	OpensIndex        uint16
	OpensFlags        uint16
	OpensToIndexTable []uint16
}

type ModuleProvides struct {
	ProvidesIndex          uint16
	ProvidesWithIndexTable []uint16
}

func readModuleAttribute(reader *ClassReader) ModuleAttribute {
	return ModuleAttribute{
		ModuleNameIndex:    reader.ReadUint16(),
		ModuleFlags:        reader.ReadUint16(),
		ModuleVersionIndex: reader.ReadUint16(),
		RequiresTable:      readRequiresTable(reader),
		ExportsTable:       readExportsTable(reader),
		OpensTable:         readOpensTable(reader),
		UsesIndexTable:     reader.readUint16s(),
		ProvidesTable:      readProvidesTable(reader),
	}
}

func readRequiresTable(reader *ClassReader) []ModuleRequires {
	return reader.readTable(func(reader *ClassReader) ModuleRequires {
		return ModuleRequires{
			RequiresIndex:        reader.ReadUint16(),
			RequiresFlags:        reader.ReadUint16(),
			RequiresVersionIndex: reader.ReadUint16(),
		}
	}).([]ModuleRequires)
}

func readExportsTable(reader *ClassReader) []ModuleExports {
	return reader.readTable(func(reader *ClassReader) ModuleExports {
		return ModuleExports{
			ExportsIndex:        reader.ReadUint16(),
			ExportsFlags:        reader.ReadUint16(),
			ExportsToIndexTable: reader.readUint16s(),
		}
	}).([]ModuleExports)
}

func readOpensTable(reader *ClassReader) []ModuleOpens {
	return reader.readTable(func(reader *ClassReader) ModuleOpens {
		return ModuleOpens{
			OpensIndex:        reader.ReadUint16(),
			OpensFlags:        reader.ReadUint16(),
			OpensToIndexTable: reader.readUint16s(),
		}
	}).([]ModuleOpens)
}

func readProvidesTable(reader *ClassReader) []ModuleProvides {
	return reader.readTable(func(reader *ClassReader) ModuleProvides {
		return ModuleProvides{
			ProvidesIndex:          reader.ReadUint16(),
			ProvidesWithIndexTable: reader.readUint16s(),
		}
	}).([]ModuleProvides)
}
