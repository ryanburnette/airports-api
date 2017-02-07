require "CSV"

class Airport
	attr_reader :id, :name, :city, :country, :iata, :icao, :latitude, :longitude,
							:elevation, :utc_offset, :_class, :timezone

  def self.new_icao(identifier, airports=nil)
    if airports
      self.new({ :by => "icao", :identifier => airport_identifier(identifier) }, airports)
    else
      self.new({ :by => "icao", :identifier => airport_identifier(identifier) })
    end
  end

  def self.new_iata(identifier, airports=nil)
    if airports
      self.new({ :by => "iata", :identifier => airport_identifier(identifier) }, airports)
    else
      self.new({ :by => "iata", :identifier => airport_identifier(identifier) })
    end
  end

  def initialize(args, airports=nil)
    airports = self.class.get_airports unless airports

    a = airports.find { |airport| airport[airport_type(args)] == args[:identifier] }.to_h

    @id         = a[:id]
    @city       = a[:city]
    @state      = a[:state]
    @country    = a[:country]
    @iata       = a[:iata]
    @icao       = a[:icao]
    @latitude   = a[:latitude]
    @longitude  = a[:longitude]
    @elevation  = a[:elevation]
    @utc_offset = a[:utc_offset]
    @_class     = a[:_class]
    @timezone   = a[:timezone]
  end

  def self.get_airports
    CSV::Converters[:blank_to_nil] = lambda do |field|
      field && field.empty? ? nil : field
    end
    CSV.new(File.open("data/airports.csv"), :headers => true, :header_converters => :symbol, :converters => [:all, :blank_to_nil])
    .map { |a| a.to_h }
  end

  def to_s
    self.iata
  end

  private

  def airport_type(args)
    by = args[:by]
    return "icao".to_sym if by == "icao"
    return "iata".to_sym if by == "iata"
    raise "Must search by a valid airport identifier type"
  end

  def self.airport_identifier(i)
    i.to_s.upcase
  end

end
