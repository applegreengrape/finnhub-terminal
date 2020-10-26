# This file was generated by GoReleaser. DO NOT EDIT.
class FinnhubTerminal < Formula
  desc "terminal dashboard for finnhub.io"
  homepage "https://github.com/applegreengrape/finnhub-terminal"
  version "0.1.8"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/applegreengrape/finnhub-terminal/releases/download/v0.1.8/finnhub-terminal_0.1.8_Darwin_x86_64.tar.gz"
    sha256 "f4421c29d514c68dc3abe44aca30e374bb8892c16ceeb36ec43ada14e45a6831"
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://github.com/applegreengrape/finnhub-terminal/releases/download/v0.1.8/finnhub-terminal_0.1.8_Linux_x86_64.tar.gz"
      sha256 "8be3fd3acbe4d00cd380dceadc7cd77e107806f503aa65336683255c1679f973"
    end
  end

  def install
    bin.install "finnhub-terminal"
  end
end