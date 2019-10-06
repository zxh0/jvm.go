package java6.string;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class StringTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(StringTest.class);
    }
    
    @Test
    public void test() {
        String s1 = "abc1";
        String s2 = "abc1";
        assertSame(s1, s2);
        
        int x = 1;
        String s3 = "abc" + x;
        assertNotSame(s1, s3);
        
        s3 = s3.intern();
        assertSame(s1, s3);
    }
    
}
