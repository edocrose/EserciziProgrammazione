public class DecorazioneElettrica extends Decorazione implements Comparable<DecorazioneElettrica> {

    private double potenza;
    private boolean accesa;

    public DecorazioneElettrica(String nome, double peso, double potenza) throws IllegalArgumentException {
        super(nome, peso);
        if (potenza <= 0) {
            throw new IllegalArgumentException("Potenza <= 0");
        }
        this.potenza = potenza;
        this.accesa = false;
    }

    public double getPotenza() {
        return potenza;
    }

    public void interruttore() {
        if (this.accesa == false) {
            this.accesa = true;
        } else {
            this.accesa = false;
        }
    }

    @Override
    public int compareTo(DecorazioneElettrica o) {
        if (this.getPeso() > o.getPeso()) {
            return 1;
        }
        if (this.getPeso() < o.getPeso()) {
            return -1;
        }
        return 0;
    }

}
