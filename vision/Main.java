// import java.util.Arrays;
import java.io.IOException;

public class Main {
  public static void main(String[] args) {
    // String file="gersedip";
    String file="aigefcop";
    // String file="savemdip";
    if (args.length > 0) {
      file = args[0];
      System.out.println("O arquivo é: " + file);
    }
    try {
      VisionFile visionFile = new VisionFile(file);
      visionFile.printInfo();

      ScanVision scanVision = new ScanVision();
      int ret = scanVision.open(file, 1, 0, null, visionFile.maxRec, visionFile.minRec, visionFile.nKeys, 0, false,
          false);
      byte[] b = new byte[visionFile.maxRec]; // Crie um array de bytes com o tamanho desejado

      // System.out.println("Retorno: " + ret);
      // System.out.println(Arrays.toString(b));
      
      long result = scanVision.start(b, 0, 0, 0, IOConstants.START_LAST);
      System.out.println("Resultado: " + result);


      for (int i = 0; i < 100; i++) {
        result = scanVision.next(b, 0, 0);
        System.out.println("Result: " + result);
        if (result == 0) {
          break;
        }
        String s = new String(b);
        System.out.println(s);
        System.out.println("-------------------------------");
      }
    } catch (IOException e) {
      e.printStackTrace();
    }

  }
}