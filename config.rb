require "byebug"

load "lib/airport.rb"
AIRPORTS = Airport.get_airports

set :layout, false

page "*.json", :content_type => "application/json"

configure :development do
end

configure :build do
end

helpers do
  def airports
    AIRPORTS
  end
end

proxy_file = "airport.json"
AIRPORTS.drop(1).each do |a|
  if a[:icao]
    proxy "/icao/#{a[:icao].to_s.downcase}.json", proxy_file, :locals => { :airport => a }
  end

  if a[:iata]
    proxy "/iata/#{a[:iata].to_s.downcase}.json", proxy_file, :locals => { :airport => a }
  end
end
