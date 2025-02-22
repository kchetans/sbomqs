# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Sbomqs < Formula
  desc ""
  homepage "https://github.com/kchetans/sbomqs"
  version "0.0.15"

  on_macos do
    url "https://github.com/kchetans/sbomqs/releases/download/v0.0.15/sbomqs-darwin-amd64"
    sha256 "5509ba48ed0a4ff09c99a94aef07187b2ef0bd4ebcaba942b0474ab5bd149204"

    def install
      bin.install "sbomqs-darwin-amd64" => "sbomqs-darwin-amd64"
    end

    if Hardware::CPU.arm?
      def caveats
        <<~EOS
          The darwin_arm64 architecture is not supported for the Sbomqs
          formula at this time. The darwin_amd64 binary may work in compatibility
          mode, but it might not be fully supported.
        EOS
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/kchetans/sbomqs/releases/download/v0.0.15/sbomqs-linux-amd64"
      sha256 "3420db9c28ff4b69ff696a7c11624c3c9a87f158ce3d0d69f8a20656f9e16aee"

      def install
        bin.install "sbomqs-linux-amd64" => "sbomqs-linux-amd64"
      end
    end
  end
end
