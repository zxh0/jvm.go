package file;

import java.io.FileDescriptor;
import java.io.FileOutputStream;
import java.io.IOException;

public class FileDescriptorTest {
    
    public static void main(String[] args) throws IOException {
        FileOutputStream fdOut = new FileOutputStream(FileDescriptor.out);
        fdOut.write("OK!\n".getBytes());
    }
    
}
