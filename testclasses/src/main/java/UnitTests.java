import junit.framework.Assert;
import org.junit.Test;
import reflection.ArrayClassTest;
import reflection.PrimitiveClassTest;
import junit.UnitTestRunner;

public class UnitTests {
    
    @Test
    public void test() {
        Assert.assertEquals(2, 1 + 1);
    }
    
    public static void main(String[] args) {
        UnitTestRunner.run(new Class<?>[] {
            UnitTests.class,
            PrimitiveClassTest.class,
            ArrayClassTest.class,
        });
    }
    
}
