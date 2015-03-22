import java7.StringTest;
import java7.cl.ClassLoaderTest;
import java7.cl.GetClassLoaderTest;
import java7.ex.ClassLibExTest;
import java7.ex.InstructionExTest;
import java7.ex.InstructionNpeTest;
import java7.file.FileIoTest;
import java7.reflection.ArrayClassTest;
import java7.reflection.ArrayGetTest;
import java7.reflection.ArraySetTest;
import java7.reflection.MethodTest;
import java7.reflection.PrimitiveClassTest;
import java7.thread.DaemonTest;
import java7.thread.MainThreadTest;
import java7.thread.SleepTest;
import java7.wrapper.DoubleTest;
import java8.InterfaceMethodTest;
import libs.junit.UnitTestRunner;

public class UnitTests {
    
    public static void main(String[] args) {
        UnitTestRunner.run(new Class<?>[] {
            ArrayClassTest.class,
            ArrayGetTest.class,
            ArraySetTest.class,
            ClassLibExTest.class,
            ClassLoaderTest.class,
            DaemonTest.class,
            DoubleTest.class,
            FileIoTest.class,
            GetClassLoaderTest.class,
            InstructionExTest.class,
            InstructionNpeTest.class,
            InterfaceMethodTest.class,
            MainThreadTest.class,
            MethodTest.class,
            PrimitiveClassTest.class,
            SleepTest.class,
            StringTest.class,
        });
    }
    
}
