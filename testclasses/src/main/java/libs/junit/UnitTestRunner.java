package libs.junit;

import org.junit.runner.JUnitCore;
import org.junit.runner.Result;
import org.junit.runner.notification.Failure;

public class UnitTestRunner {
    
    // todo
    public static void run(Class<?>... tests) {
        Result result = JUnitCore.runClasses(tests);
        System.out.println("total: " + result.getRunCount());
        System.out.println("failures: " + result.getFailureCount());
        
        if (!result.wasSuccessful()) {
            for (Failure f : result.getFailures()) {
                System.out.println(f);
                //f.getException().printStackTrace();
            }
        }
    }
    
}
