import helper.UnitTestRunner;
import jvm.lambda.InterfaceMethodTest;

public class UnitTests {

    public static void main(String[] args) {
        UnitTestRunner.run(
            stdlib.basic.cl.ClassLoaderTest.class,
            stdlib.basic.cl.GetClassLoaderTest.class,
            jvm.ex.ClassLibExTest.class,
            jvm.ex.InstructionExTest.class,
            jvm.ex.InstructionNpeTest.class,
            jvm.field.ConstantStaticFieldsTest.class,
            jvm.field.FieldAccessTest.class,
            jvm.field.FieldsTest.class,
            stdlib.nio.io.file.FileIoTest.class,
            stdlib.nio.io.file.RandomAccessFileTest.class,
            stdlib.basic.reflection.ArrayClassTest.class,
            stdlib.basic.reflection.ArrayGetTest.class,
            stdlib.basic.reflection.ArraySetTest.class,
            stdlib.basic.reflection.GenericTest.class,
            stdlib.basic.reflection.MethodTest.class,
            stdlib.basic.reflection.PrimitiveClassTest.class,
            stdlib.basic.string.StringTest.class,
            stdlib.basic.string.Mutf8Test.class,
            stdlib.basic.sunmisc.UnsafeMemoryTest.class,
            stdlib.basic.sunmisc.UnsafeObjectTest.class,
            stdlib.basic.thread.DaemonTest.class,
            stdlib.basic.thread.MainThreadTest.class,
            stdlib.basic.thread.SleepTest.class,
            stdlib.basic.wrapper.DoubleTest.class,
            //jvm.lambda.InterfaceDefaultMethodTest.class,
            InterfaceMethodTest.class,
            jls8.ch12.Eg12_4_1_1.class,
            jls8.ch12.Eg12_4_1_2.class,
            jls8.ch12.Eg12_4_1_3.class,
            jls8.ch12.Eg12_5_2.class,
            thirdparty.gson.GsonTest.class
        );
    }

}
