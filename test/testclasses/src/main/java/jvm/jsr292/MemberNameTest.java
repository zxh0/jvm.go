package jvm.jsr292;

import java.lang.reflect.Constructor;
import java.lang.reflect.Field;
import java.lang.reflect.Method;

import static helper.MyAssert.assertEquals;

public class MemberNameTest {

    public static void main(String[] args) throws Exception {
        Method mainMethod = MemberNameTest.class.getDeclaredMethod("main", String[].class);

        Class<?> mnClass = Class.forName("java.lang.invoke.MemberName");
        Constructor<?> mnConstructor = mnClass.getConstructor(Method.class);
        mnConstructor.setAccessible(true);
        Field flags = mnClass.getDeclaredField("flags");
        flags.setAccessible(true);
        Method getFieldType = mnClass.getDeclaredMethod("getMethodOrFieldType");
        getFieldType.setAccessible(true);
        Method getRefKind = mnClass.getDeclaredMethod("getReferenceKind");
        getRefKind.setAccessible(true);

        Object mn = mnConstructor.newInstance(mainMethod);
        assertEquals(100728841, flags.get(mn));
        assertEquals((byte)6, getRefKind.invoke(mn));
        assertEquals("(String[])void", getFieldType.invoke(mn).toString());

        System.out.println("OK!");
    }

}
