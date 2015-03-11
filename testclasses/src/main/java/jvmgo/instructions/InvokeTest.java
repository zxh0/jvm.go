package jvmgo.instructions;

import java.lang.reflect.Constructor;
import java.lang.reflect.Field;
import java.lang.reflect.Method;
import java.util.ArrayList;
import java.util.List;
import jvmgo.UnitTestRunner;
import org.junit.Test;
import org.junit.runners.model.FrameworkField;
import org.junit.runners.model.FrameworkMember;
import org.junit.runners.model.FrameworkMethod;
import static org.junit.Assert.*;

@SuppressWarnings("serial")
public class InvokeTest extends ArrayList<String> {
    
    public int x;

    public static void main(String[] args) {
        UnitTestRunner.run(InvokeTest.class);
    }
    
    @Override
    public String get(int index) {
        return "InvokeTest";
    }
    
    public String getFromSuper(int index) {
        return super.get(index);
    }
    
    @Test
    public void invoke() {
        InvokeTest test = new InvokeTest();
        test.add("ArrayList");
        
        // invokevirtual
        ArrayList<String> arrayList = test;
        assertEquals("InvokeTest", arrayList.get(0));
        
        // invokespecial
        assertEquals("ArrayList", test.getFromSuper(0));
        
        // invokeinterface
        List<String> list = test;
        assertEquals("InvokeTest", list.get(0));
    }
    
    @Test
    public void invokeVirtual() throws Exception {
        Constructor<FrameworkField> c = FrameworkField.class.getDeclaredConstructor(Field.class);
        c.setAccessible(true);
        Field f = InvokeTest.class.getField("x");
        FrameworkField ff = c.newInstance(f);
        
        Method m = InvokeTest.class.getMethod("invokeVirtual");
        FrameworkMethod fm = new FrameworkMethod(m);
        
        FrameworkMember<?> a = ff;
        FrameworkMember<?> b = fm;
        
        a.getAnnotations();
        b.getAnnotations();
    }
    
}
