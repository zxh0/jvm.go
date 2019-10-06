package java6.file;

import java.io.FileInputStream;
import java.io.FileNotFoundException;
import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class FileIoTest {
    
    public static void main(String[] args) throws Exception {
        UnitTestRunner.run(FileIoTest.class);
    }
    
    @Test
    public void fileNotFoundException() {
        try {
            FileInputStream fis = new FileInputStream("a/b/foo.txt");
        } catch (FileNotFoundException e) {
            assertEquals("a/b/foo.txt (No such file or directory)", e.getMessage());
        }
    }
    
}
