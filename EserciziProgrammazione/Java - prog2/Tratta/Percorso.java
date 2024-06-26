import java.util.ArrayList;
import java.util.Iterator;

public class Percorso implements Iterable<Tratta>, Comparable<Percorso> {
    // OVERVIEW: modella un percorso che è un insieme di tratte
    private ArrayList<Tratta> tratte = new ArrayList<>();

    public void add(Tratta t) throws TrattaNonValidaException {
        // MODIFIES: this
        // EFFECTS: aggiunge t all'array
        // se l'origine di t != dalla destinazione dell'ultima tratta del percorso
        // lancia TrattaNonValidaException
        if (tratte.size() != 0) {
            if (!(t.getOrigine().equals(tratte.get(tratte.size() - 1).getDestinazione()))) {
                throw new TrattaNonValidaException("Tratta non valida");
            }
        }
        tratte.add(t);
    }

    public double calcolaDurataTotale() {
        // EFFECTS: calcola la durata totale del percorso con la somma delle durate
        // delle tratte
        double d = 0;
        for (Tratta t : tratte) {
            d += t.calcolaDurata();
        }
        return d;
    }

    public double calcolaCO2Totale() {
        // EFFECTS: calcola la quantità di CO2 totale del percorso come somma delle
        // quantità di ogni tratta
        double c = 0;
        for (Tratta t : tratte) {
            c += t.calcolaCO2();
        }
        return c;
    }

    @Override
    public Iterator<Tratta> iterator() {
        return new Iterator<Tratta>() {
            Iterator<Tratta> i = tratte.iterator();

            @Override
            public boolean hasNext() {
                return i.hasNext();
            }

            @Override
            public Tratta next() {
                return i.next();
            }

        };
    }

    @Override
    public int compareTo(Percorso o) {
        if (this.calcolaDurataTotale() > o.calcolaDurataTotale()) {
            return 1;
        }
        if (this.calcolaDurataTotale() < o.calcolaDurataTotale()) {
            return -1;
        }
        return 0;
    }

    @Override
    public String toString() {
        return "Percorso (durata: " + this.calcolaDurataTotale() + ", co2: " + this.calcolaCO2Totale() + ")";
    }

}
