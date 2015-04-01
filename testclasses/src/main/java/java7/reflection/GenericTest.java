package java7.reflection;

import java.util.ArrayList;
import java.util.List;
import org.junit.Test;

public class GenericTest {
    
    private static final List<String> list = new ArrayList<>();
    
    @Test
    public void fieldGenericType() {
        
    }
    
    public static void main(String[] args) throws Exception {
        System.out.println(GenericTest.class.getDeclaredField("list").getGenericType());
    }
    
}
