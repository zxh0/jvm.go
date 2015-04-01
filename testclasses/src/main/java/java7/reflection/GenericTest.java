package java7.reflection;

import java.lang.reflect.Field;
import java.util.ArrayList;
import java.util.List;
import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class GenericTest {
    
    private static class GenericClass<String> {
    
    }
    
    private static final String str = "abc";
    private static final List<String> list = new ArrayList<>();
    
    @Test
    public void typeParameter() {
        assertEquals(1, GenericClass.class.getTypeParameters().length);
        assertEquals("String", GenericClass.class.getTypeParameters()[0].toString());
    }
    
    @Test
    public void fieldGenericType() throws NoSuchFieldException {
        assertEquals(String.class,
                GenericTest.class.getDeclaredField("str").getGenericType());
        
        Field listField = GenericTest.class.getDeclaredField("list");
        assertEquals("list", listField.getName());
        assertEquals(List.class, listField.getType());
        assertEquals(26, listField.getModifiers());
//        System.out.println(f.getGenericType());
    }
    
    public static void main(String[] args) throws Exception {
        UnitTestRunner.run(GenericTest.class);
    }
    
}
