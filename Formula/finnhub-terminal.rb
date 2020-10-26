# This file was generated by GoReleaser. DO NOT EDIT.
class FinnhubTerminal < Formula
  desc "terminal dashboard for finnhub.io"
  homepage "https://github.com/applegreengrape/finnhub-terminal"
  version "0.1.9"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/applegreengrape/finnhub-terminal/releases/download/v0.1.9/finnhub-terminal_0.1.9_darwin_amd64.zip"
    sha256 "2a6c9346ec4557d17ddd9898c24884800b6157a0941c0d1b411341352b779164"
  elsif OS.linux?
  end

  def install
    bin.install "golang-cross"
  end
end
