require "yaml"
require "fileutils"
require "byebug"
load "lib/airport.rb"

desc "Deploy"
task :deploy do
  sh "middleman build"
  sh "aws s3 sync build/ s3://airports.api.faralmanac.com/ --delete --acl public-read"
end

def slugify(text)
  text.to_s.downcase.strip.gsub(' ', '-').gsub(/[^\w-]/, '')
end

def write_airport(dir, filename, data)
  path = "./#{dir}/#{filename}.yml"
  FileUtils.touch(path)
  File.open(path, "w") { |f| f.write data.to_yaml }
end

task :convert do
  Airport.get_airports.each do |airport|
    airport.delete(:id)
    airport = Hash[airport.map{ |k, v| [k.to_s, v] }]

    if airport["icao"] == nil && airport["iata"] == nil
      write_airport("data.incomplete", slugify(airport["airport_name"]), airport)
    elsif airport["icao"] == nil && airport["iata"]
      write_airport("data", "iata-"+airport["iata"].to_s.downcase, airport)
    else
      write_airport("data", airport["icao"].to_s.downcase, airport)
    end
  end
end
