package classpath

type ClassPath struct {
    entries []ClassPathEntry
}

type ClassPathEntry interface {

}

type ClassPathDirEntry struct {
    // todo
}

type ClassPathJarEntry struct {
    // todo
}

func Parse(cpOption string) (*ClassPath) {
    // todo
    return nil
}
