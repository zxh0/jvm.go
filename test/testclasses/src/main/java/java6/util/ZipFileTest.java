package java6.util;

import java.util.Collections;
import java.util.zip.ZipEntry;
import java.util.zip.ZipFile;

public class ZipFileTest {
    
    public static void main(String[] args) throws Exception {
        ZipFile zf = new ZipFile("/Users/zxh/Work/GitHub/jvm.go/testclasses/build/libs/testclasses.jar");
        
        for (ZipEntry x : Collections.list(zf.entries())) {
            System.out.println(x);
        }
    }
    
}
