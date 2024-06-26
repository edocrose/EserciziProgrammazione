import java.util.ArrayList;
import java.util.Iterator;

public class AlberoNatalizio implements Iterable<Decorazione> {
    private double potenzaMax;
    private double caricoMax;
    private ArrayList<Decorazione> decorazioni = new ArrayList<>();

    public AlberoNatalizio(double caricoMax, double potenzaMax) {
        if (caricoMax <= 0) {
            throw new IllegalArgumentException("Carico massimo <= 0");
        }
        if (potenzaMax <= 0) {
            throw new IllegalArgumentException("Potenza massima <= 0");
        }
        this.caricoMax = caricoMax;
        this.potenzaMax = potenzaMax;
        assert repOk();
    }

    public double getCaricoMax() {
        return this.caricoMax;
    }

    public double getPotenzaMax() {
        return this.potenzaMax;
    }

    public void add(Decorazione d) throws WeightReachedException, TopperExistsException {
        if (d == null) {
            throw new IllegalArgumentException("Decorazione nulla");
        }
        if (this.pesoCorrente() + d.getPeso() > this.getCaricoMax()) {
            throw new WeightReachedException("Carico superato");
        }
        if (d.getClass().equals(Puntale.class)) {
            if (contaPuntale() > 0) {
                throw new TopperExistsException("Puntale già esistente");
            }
        }
        decorazioni.add(d);
        assert repOk();
    }

    public void accendi() {
        ArrayList<DecorazioneElettrica> de = new ArrayList<>();
        for (Decorazione d : decorazioni) {
            if (d.getClass().equals(DecorazioneElettrica.class)) {
                de.add((DecorazioneElettrica) d);
            }
        }
        de.sort(null);
        for (DecorazioneElettrica dc : de) {
            if (this.potenzaCorrente() + dc.getPotenza() <= this.potenzaMax) {
                dc.interruttore();
            }
        }
        assert repOk();
    }

    public double pesoCorrente() {
        return 0;
    }

    public double potenzaCorrente() {
        return 0;
    }

    public int contaPuntale() {
        int c = 0;
        for (Decorazione d : decorazioni) {
            if (d.getClass().equals(Puntale.class)) {
                c++;
            }
        }
        return c;
    }

    public boolean repOk() {
        for (Decorazione d : decorazioni) {
            if (d == null) {
                return false;
            }
        }
        if (this.contaPuntale() > 1) {
            return false;
        }
        if (this.pesoCorrente() > this.caricoMax) {
            return false;
        }
        if (this.potenzaCorrente() > this.potenzaMax) {
            return false;
        }
        return true;
    }

    @Override
    public Iterator<Decorazione> iterator() {
        return new Iterator<Decorazione>() {
            Iterator<Decorazione> i = decorazioni.iterator();

            @Override
            public boolean hasNext() {
                return i.hasNext();
            }

            @Override
            public Decorazione next() {
                return i.next();
            }

        };
    }

}
