public class TrattaBus extends Tratta {

    private double co2ora;

    public TrattaBus(String origine, String destinazione, double lung, double vel, double co2ora)
            throws IllegalArgumentException {
        // MODIFIES: this
        // EFFECTS: inizializza una tratta in aereo con gli attributi di Tratta + un
        // altro
        // se il nuovo parametro non è corretto lancia IllegalArgumentException
        super(origine, destinazione, lung, vel);
        if (co2ora <= 0) {
            throw new IllegalArgumentException("Quantità di CO2/h <= 0");
        }
        this.co2ora = co2ora;
    }

    @Override
    public double calcolaCO2() {
        // EFFECTS: restiuisce la quantità di CO2 prodotta come durata * co2ora
        return this.calcolaDurata() * this.co2ora;
    }

    @Override
    public String toString() {
        return "(Bus) Tratta da " + this.getOrigine() + " a " + this.getDestinazione() + "; lunghezza: "
                + this.getLung() + ", velocità: " + this.getVel() + ", CO2/h: " + this.co2ora;
    }

}
