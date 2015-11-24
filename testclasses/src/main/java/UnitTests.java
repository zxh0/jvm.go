import libs.junit.UnitTestRunner;

public class UnitTests {

    public static void main(String[] args) {
        UnitTestRunner.run(
            java6.cl.ClassLoaderTest.class,
            java6.cl.GetClassLoaderTest.class,
            java6.ex.ClassLibExTest.class,
            java6.ex.InstructionExTest.class,
            java6.ex.InstructionNpeTest.class,
            java6.field.ConstantStaticFieldsTest.class,
            java6.field.FieldAccessTest.class,
            java6.field.FieldsTest.class,
            java6.file.FileIoTest.class,
            java6.file.RandomAccessFileTest.class,
            java6.reflection.ArrayClassTest.class,
            java6.reflection.ArrayGetTest.class,
            java6.reflection.ArraySetTest.class,
            java6.reflection.GenericTest.class,
            java6.reflection.MethodTest.class,
            java6.reflection.PrimitiveClassTest.class,
            java6.string.StringTest.class,
            java6.string.Mutf8Test.class,
            java6.sunmisc.UnsafeMemoryTest.class,
            java6.sunmisc.UnsafeObjectTest.class,
            java6.thread.DaemonTest.class,
            java6.thread.MainThreadTest.class,
            java6.thread.SleepTest.class,
            java6.wrapper.DoubleTest.class,
            //java8.InterfaceDefaultMethodTest.class,
            java8.InterfaceMethodTest.class,
            jls8.ch12.Eg12_4_1_1.class,
            jls8.ch12.Eg12_4_1_2.class,
            jls8.ch12.Eg12_4_1_3.class,
            jls8.ch12.Eg12_5_2.class,
            libs.gson.GsonTest.class
        );
    }

}
