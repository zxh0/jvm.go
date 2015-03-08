package file;

import java.io.FileInputStream;

public class FileIoTest {
    
    public static void main(String[] args) throws Exception {
        FileInputStream fis = new FileInputStream("foo.txt");
        //fis.close();
    }
    
}
