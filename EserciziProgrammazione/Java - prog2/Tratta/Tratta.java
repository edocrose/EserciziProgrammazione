abstract class Tratta {
    // OVERVIEW: definisce una tratta tra due luoghi

    private String origine, destinazione;
    private double lung, vel;

    public Tratta(String origine, String destinazione, double lung, double vel) throws IllegalArgumentException {
        // MODIFIES: this
        // EFFECTS: inizializza una tratta con origine, destinazione, ecc.
        // se i parametri non sono corretti lancia IllegalArgumentException

        if (origine.equals("") || origine == null) {
            throw new IllegalArgumentException("Origine nulla o vuota");
        }
        if (destinazione.equals("") || destinazione == null) {
            throw new IllegalArgumentException("Destinazione nulla o vuota");
        }
        if (lung <= 0) {
            throw new IllegalArgumentException("Lunghezza <= 0");
        }
        if (vel <= 0) {
            throw new IllegalArgumentException("Lunghezza <= 0");
        }

        this.origine = origine;
        this.destinazione = destinazione;
        this.lung = lung;
        this.vel = vel;

        assert repOk();
    }

    public String getOrigine() {
        return this.origine;
    }

    public String getDestinazione() {
        return this.destinazione;
    }

    public double getLung() {
        return this.lung;
    }

    public double getVel() {
        return this.vel;
    }

    public double calcolaDurata() {
        // EFFECTS: calcola in ore la durata della tratta
        return this.lung / this.vel;
    }

    public abstract double calcolaCO2();
    // EFFECTS: calcola la quantità di CO2 prodotta durante la tratta

    public boolean repOk() {
        if (this.origine.equals("") || this.origine == null) {
            return false;
        }
        if (this.destinazione.equals("") || this.destinazione == null) {
            return false;
        }
        if (this.lung <= 0) {
            return false;
        }
        if (this.vel <= 0) {
            return false;
        }
        return true;
    }
}
