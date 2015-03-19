package java7.file;

import libs.junit.UnitTestRunner;
import org.junit.Test;

import java.io.IOException;
import java.io.RandomAccessFile;

/**
 * Created with IntelliJ IDEA.
 * User: beyond
 * Email: beyondblog@outlook.com
 * Date: 15/3/19
 * Time: 上午7:01
 * Description:
 * java/io/RandomAccessFile#open(Ljava/lang/String;I)V
 * java/io/RandomAccessFile#read0()I
 * java/io/RandomAccessFile#readBytes([BII)I
 * java/io/RandomAccessFile#write0(I)V
 * java/io/RandomAccessFile#writeBytes([BII)V
 * java/io/RandomAccessFile#getFilePointer()J
 * java/io/RandomAccessFile#seek0(J)V
 * java/io/RandomAccessFile#length()J
 * java/io/RandomAccessFile#setLength(J)V
 * java/io/RandomAccessFile.initIDs()V
 * java/io/RandomAccessFile#close0()V
 */
public class RandomAccessFileTest {
    public static void main(String[] args) throws Exception {
        UnitTestRunner.run(RandomAccessFileTest.class);
    }

    @Test
    public void writeTest() throws Exception {
        RandomAccessFile randomAccessFile = null;
        try {
            randomAccessFile = new RandomAccessFile("RandomAccessFileWriteTest", "rw");
            //write data
            randomAccessFile.writeBytes("hello");
            randomAccessFile.writeInt(54);
            randomAccessFile.writeBytes("world");
            randomAccessFile.writeInt(57);
        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            randomAccessFile.close();
        }
    }
}
