package libs.gson;

//import com.google.gson.Gson;
import libs.junit.UnitTestRunner;
import org.junit.Test;
//import static org.junit.Assert.*;

public class GsonTest {
    
    public int x;
    public float y;
    public String z;
    
    public static void main(String[] args) {
        UnitTestRunner.run(GsonTest.class);
    }
    
    @Test
    public void gson() {
        GsonTest obj = new GsonTest();
        obj.x = 100;
        obj.y = 3.14f;
        obj.z = "hello";
//        assertEquals("{\"x\":100,\"y\":3.14,\"z\":\"hello\"}", new Gson().toJson(obj));
    }
    
}
