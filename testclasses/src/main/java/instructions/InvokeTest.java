package instructions;

import java.util.ArrayList;
import java.util.List;

@SuppressWarnings("serial")
public class InvokeTest extends ArrayList<String> {

    @Override
    public String get(int index) {
        return "InvokeTest";
    }
    
    public String getFromSuper(int index) {
        return super.get(index);
    }
    
    public static void main(String[] args) {
        InvokeTest test = new InvokeTest();
        test.add("ArrayList");
        
        // invokevirtual
        ArrayList<String> arrayList = test;
        System.out.println(arrayList.get(0)); // InvokeTest
        
        // invokespecial
        System.out.println(test.getFromSuper(0)); // ArrayList
        
        // invokeinterface
        List<String> list = test;
        System.out.println(list.get(0)); // InvokeTest
    }
    
}
