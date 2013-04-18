#
# Build helpers for working with Project Slartibartfast
#

PACKAGES = %w(
  window
  slartibartfast
)

LIBRARIES = %w(
  github.com/go-gl/gl
  github.com/go-gl/glfw
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

namespace :update do
  desc "Update all installed libraries"
  task :all => LIBRARIES.map {|l| l.split("\/").last }

  LIBRARIES.each do |library|
    desc "Update the library [#{library}]"
    task library.split("\/").last do
      sh "go get #{library}"
    end
  end
end
