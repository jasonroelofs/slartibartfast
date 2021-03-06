#
# Build helpers for working with Project Slartibartfast
#

PACKAGES = %w(
	core
  input
  events
  configs
  components
  slartibartfast
  platform
  behaviors
  math3d
  render
  freeimage
  factories
  volume
  util
  game
)

# These are git submodules
# Update with: git submodule update --remote
LIBRARIES = %w(
  github.com/go-gl/gl
  github.com/go-gl/glu
  github.com/go-gl/glfw3
  github.com/stretchr/testify
)

EXTERNAL_LIBS = %w(
  glfw3
  freeimage
)

def shell(command)
  begin
    sh command
  rescue RuntimeError
  end
end

task :default => "test"

desc "Run the application itself"
task :run => ["clean:bin", "build:all"] do
  full_dir = File.expand_path("../", __FILE__)
  shell "#{full_dir}/bin/slartibartfast"
end

task :clean => "clean:bin"
namespace :clean do
  desc "Clean up the bin directory"
  task :bin do
    full_dir = File.expand_path("../", __FILE__)
    rm_f "#{full_dir}/bin/slartibartfast"
  end
end

task :build => "build:all"
namespace :build do
  desc "Build and install all packages"
  task :all do
    shell "go install -v #{PACKAGES.join(" ")}"
  end

  PACKAGES.each do |package|
    desc "Build and install package [#{package}]"
    task package do
      shell "go install -v #{package}"
    end
  end
end

task :test => "test:all"
namespace :test do
	desc "Run all the tests. Also runnable as just 'test'"
	task :all => ["build:all", PACKAGES].flatten

	PACKAGES.each do |package|
		desc "Run tests for [#{package}]"
		task package do
			shell "go test #{package}"
		end
	end
end
