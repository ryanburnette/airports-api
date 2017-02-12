set :layout, false
page "*.json", :content_type => "application/json"

configure :development do
end

configure :build do
end

helpers do
end

proxy_file = "airport.json"
ignore proxy_file
def airports
  Dir["./data/*.yml"].collect { |p| p.gsub("./data/", "").gsub(".yml", "") }
end
airports.each do |f|
  a = data.try(:f)
  if a[:icao]
    proxy "/icao/#{a[:icao].to_s.downcase}.json", proxy_file, :locals => { :airport => a }
  end
  if a[:iata]
    proxy "/iata/#{a[:iata].to_s.downcase}.json", proxy_file, :locals => { :airport => a }
  end
end
