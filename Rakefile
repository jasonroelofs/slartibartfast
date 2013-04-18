#
# Build helpers for working with Project Slartibartfast
#

PACKAGES = %w(
  window
  slartibartfast
)

task :default => "run"

desc "Run the application itself"
task :run => "build:all" do
  full_dir = File.expand_path("../", __FILE__)
  sh "#{full_dir}/bin/slartibartfast"
end

namespace :build do
  desc "Build and install all packages"
  task :all do
    sh "go install -v #{PACKAGES.join(" ")}"
  end

  PACKAGES.each do |package|
    desc "Build and install package [#{package}]"
    task package do
      sh "go install -v #{package}"
    end
  end
end
