package jvmgo.instructions;

import java.util.ArrayList;
import java.util.List;
import jvmgo.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

@SuppressWarnings("serial")
public class InvokeTest extends ArrayList<String> {

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
    
}
