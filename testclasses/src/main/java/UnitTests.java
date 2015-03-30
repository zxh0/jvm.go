import java7.StringTest;
import java7.cl.ClassLoaderTest;
import java7.cl.GetClassLoaderTest;
import java7.ex.ClassLibExTest;
import java7.ex.InstructionExTest;
import java7.ex.InstructionNpeTest;
import java7.field.ConstantVariablesTest;
import java7.field.FieldAccessTest;
import java7.field.FieldsTest;
import java7.file.FileIoTest;
import java7.reflection.ArrayClassTest;
import java7.reflection.ArrayGetTest;
import java7.reflection.ArraySetTest;
import java7.reflection.MethodTest;
import java7.reflection.PrimitiveClassTest;
import java7.sunmisc.UnsafeMemoryTest;
import java7.sunmisc.UnsafeObjectTest;
import java7.thread.DaemonTest;
import java7.thread.MainThreadTest;
import java7.thread.SleepTest;
import java7.wrapper.DoubleTest;
import java8.InterfaceMethodTest;
import jls8.ch12.Eg12_4_1_1;
import jls8.ch12.Eg12_4_1_2;
import jls8.ch12.Eg12_4_1_3;
import libs.junit.UnitTestRunner;

public class UnitTests {
    
    public static void main(String[] args) {
        UnitTestRunner.run(new Class<?>[] {
            ArrayClassTest.class,
            ArrayGetTest.class,
            ArraySetTest.class,
            ClassLibExTest.class,
            ClassLoaderTest.class,
            ConstantVariablesTest.class,
            DaemonTest.class,
            DoubleTest.class,
            Eg12_4_1_1.class,
            Eg12_4_1_2.class,
            Eg12_4_1_3.class,
            FieldAccessTest.class,
            FieldsTest.class,
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
            UnsafeMemoryTest.class,
            UnsafeObjectTest.class,
        });
    }
    
}
