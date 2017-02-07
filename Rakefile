desc "Deploy"
task :deploy do
  sh "middleman build"
  sh "aws s3 sync build/ s3://airports.api.faralmanac.com/ --delete --acl public-read"
end
