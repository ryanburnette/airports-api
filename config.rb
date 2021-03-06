load "lib/load_airports.rb"

set :layout, false
page "*.json", :content_type => "application/json"

configure :development do
end

configure :build do
end

helpers do
  def airports
    load_airports
  end
end

proxy_file = "airport.json"
ignore proxy_file
load_airports.each do |a|
  a.delete(:id)
  if a[:icao]
    proxy "/icao/#{a[:icao].to_s.downcase}.json", proxy_file, :locals => { :airport => a }
  end
  if a[:iata]
    proxy "/iata/#{a[:iata].to_s.downcase}.json", proxy_file, :locals => { :airport => a }
  end
end
