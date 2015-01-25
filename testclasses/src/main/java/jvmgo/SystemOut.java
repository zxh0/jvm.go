package jvmgo;

import java.io.FileNotFoundException;
import java.io.PrintStream;

/**
 * To let HelloWord work!
 */
public class SystemOut extends PrintStream {

    public SystemOut() throws FileNotFoundException{
        super((String) null);
    }

    @Override
    public native void println(String x);

}
