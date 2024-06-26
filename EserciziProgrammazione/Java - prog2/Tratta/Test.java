import java.util.ArrayList;
import java.util.Scanner;

public class Test {
    public static void main(String[] args) {
        ArrayList<Percorso> percorsi = new ArrayList<>();
        Scanner s = new Scanner(System.in);
        Percorso p = new Percorso();
        percorsi.add(p);
        while (s.hasNext()) {
            String tipoTratta = s.next();
            try {
                switch (tipoTratta) {
                    case "TrattaTreno":
                        p.add(new TrattaTreno(s.next(), s.next(), s.nextDouble(), s.nextDouble(), s.nextDouble()));
                        break;
                    case "TrattaBus":
                        p.add(new TrattaBus(s.next(), s.next(), s.nextDouble(), s.nextDouble(), s.nextDouble()));
                        break;
                    case "TrattaAereo":
                        p.add(new TrattaAereo(s.next(), s.next(), s.nextDouble(), s.nextDouble(), s.nextDouble(),
                                s.nextDouble()));
                        break;
                    case "-":
                        p = new Percorso();
                        percorsi.add(p);
                        break;
                }
            } catch (TrattaNonValidaException e) {
                System.out.println(e.getMessage());
            }
        }

        percorsi.sort(null);
        for (Percorso perc : percorsi) {
            System.out.println();
            System.out.println("\t" + perc);
            for (Tratta t : perc) {
                System.out.println("\t\t" + t);
            }
        }

    }

}
