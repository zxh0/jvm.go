import java6.StringTest;
import java6.cl.ClassLoaderTest;
import java6.cl.GetClassLoaderTest;
import java6.ex.ClassLibExTest;
import java6.ex.InstructionExTest;
import java6.ex.InstructionNpeTest;
import java6.field.ConstantStaticFieldsTest;
import java6.field.FieldAccessTest;
import java6.field.FieldsTest;
import java6.file.FileIoTest;
import java6.file.RandomAccessFileTest;
import java6.reflection.ArrayClassTest;
import java6.reflection.ArrayGetTest;
import java6.reflection.ArraySetTest;
import java6.reflection.GenericTest;
import java6.reflection.MethodTest;
import java6.reflection.PrimitiveClassTest;
import java6.sunmisc.UnsafeMemoryTest;
import java6.sunmisc.UnsafeObjectTest;
import java6.thread.DaemonTest;
import java6.thread.MainThreadTest;
import java6.thread.SleepTest;
import java6.wrapper.DoubleTest;
import java8.InterfaceDefaultMethodTest;
import java8.InterfaceMethodTest;
import jls8.ch12.Eg12_4_1_1;
import jls8.ch12.Eg12_4_1_2;
import jls8.ch12.Eg12_4_1_3;
import jls8.ch12.Eg12_5_2;
import libs.gson.GsonTest;
import libs.junit.UnitTestRunner;

public class UnitTests {

    public static void main(String[] args) {
        UnitTestRunner.run(new Class<?>[]{
                ArrayClassTest.class,
                ArrayGetTest.class,
                ArraySetTest.class,
                ClassLibExTest.class,
                ClassLoaderTest.class,
                ConstantStaticFieldsTest.class,
                DaemonTest.class,
                DoubleTest.class,
                Eg12_4_1_1.class,
                Eg12_4_1_2.class,
                Eg12_4_1_3.class,
                Eg12_5_2.class,
                FieldAccessTest.class,
                FieldsTest.class,
                FileIoTest.class,
                GenericTest.class,
                GetClassLoaderTest.class,
                InstructionExTest.class,
                InstructionNpeTest.class,
                InterfaceDefaultMethodTest.class,
                InterfaceMethodTest.class,
                MainThreadTest.class,
                MethodTest.class,
                PrimitiveClassTest.class,
                SleepTest.class,
                StringTest.class,
                UnsafeMemoryTest.class,
                UnsafeObjectTest.class,
                RandomAccessFileTest.class,
                GsonTest.class,
        });
    }

}
