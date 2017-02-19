require "CSV"

def load_airports
  airports_file = "airport-data/airports.csv"
  CSV::Converters[:blank_to_nil] = lambda do |field|
    field && field.empty? ? nil : field
  end
  CSV.new(File.open(airports_file), :headers => true, :header_converters => :symbol, :converters => [:all, :blank_to_nil])
  .map { |a| a.to_h }
end
