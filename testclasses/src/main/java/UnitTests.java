import junit.framework.Assert;
import org.junit.Test;
import org.junit.runner.JUnitCore;
import org.junit.runner.Result;
import org.junit.runner.notification.Failure;

public class UnitTests {
    
    @Test
    public void test() {
        Assert.assertEquals(2, 1 + 1);
    }
    
    public static void main(String[] args) {
        Result result = JUnitCore.runClasses(UnitTests.class);
        if (result.getFailureCount() > 0) {
            for (Failure f : result.getFailures()) {
                System.out.println(f);
            }
        } else {
            System.out.println("OK!");
        }
    }
    
}
