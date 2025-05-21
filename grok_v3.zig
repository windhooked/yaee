const std = @import("std");

const Enigma = struct {
    rotors: [3]Rotor,
    reflector: [26]u8,
    plugboard: [26]u8,
    allocator: std.mem.Allocator,

    const Rotor = struct {
        wiring: [26]u8,
        position: u8,
        notch: u8,
    };

    fn init(allocator: std.mem.Allocator) !Enigma {
        // Rotor wirings (Enigma I: Rotor I, II, III)
        const rotor1 = [_]u8{ 4, 10, 12, 5, 11, 6, 3, 16, 21, 25, 13, 19, 14, 22, 24, 7, 23, 20, 18, 15, 0, 8, 1, 17, 2, 9 };
        const rotor2 = [_]u8{ 0, 9, 3, 10, 18, 8, 17, 20, 23, 1, 11, 7, 22, 19, 12, 2, 16, 6, 25, 13, 15, 24, 5, 21, 14, 4 };
        const rotor3 = [_]u8{ 1, 3, 5, 7, 9, 11, 2, 15, 17, 19, 23, 21, 25, 13, 24, 4, 8, 22, 6, 0, 10, 12, 20, 18, 16, 14 };

        // Reflector B
        const reflector = [_]u8{ 24, 17, 20, 7, 16, 18, 11, 3, 15, 23, 13, 6, 14, 10, 12, 8, 4, 1, 5, 25, 2, 22, 21, 9, 0, 19 };

        // Plugboard with swaps: A<->B, C<->D
        var plugboard = [_]u8{0} ** 26;
        for (0..26) |i| plugboard[i] = @intCast(i);
        plugboard[0] = 1; plugboard[1] = 0; // A<->B
        plugboard[2] = 3; plugboard[3] = 2; // C<->D

        return Enigma{
            .rotors = [_]Rotor{
                .{ .wiring = rotor1, .position = 0, .notch = 16 }, // Rotor I, notch at Q
                .{ .wiring = rotor2, .position = 0, .notch = 4 },  // Rotor II, notch at E
                .{ .wiring = rotor3, .position = 0, .notch = 21 }, // Rotor III, notch at V
            },
            .reflector = reflector,
            .plugboard = plugboard,
            .allocator = allocator,
        };
    }

    fn setPositions(self: *Enigma, pos: []const u8) void {
        if (pos.len != 3) return;
        for (0..3) |i| {
            self.rotors[i].position = pos[i] - 'A';
        }
    }

    fn stepRotors(self: *Enigma) void {
        // Right rotor always steps
        self.rotors[2].position = (self.rotors[2].position + 1) % 26;

        // Middle rotor steps if right rotor is at notch (double-stepping)
        if (self.rotors[2].position == self.rotors[2].notch) {
            self.rotors[1].position = (self.rotors[1].position + 1) % 26;
        }

        // Left rotor steps if middle rotor is at notch
        if (self.rotors[1].position == self.rotors[1].notch) {
            self.rotors[0].position = (self.rotors[0].position + 1) % 26;
        }
    }

    fn encryptChar(self: *Enigma, c: u8) u8 {
        if (c >= 26) return c; // Non-alphabetic, return unchanged

        // Step rotors before processing
        self.stepRotors();

        // Forward pass: plugboard -> rotors (right to left)
        var current = self.plugboard[c];
        for (self.rotors) |rotor| {
            const offset = rotor.position;
            current = (rotor.wiring[(current + offset) % 26] + 26 - offset) % 26;
        }

        // Reflector
        current = self.reflector[current];

        // Backward pass: rotors (left to right)
        var i: usize = self.rotors.len;
        while (i > 0) : (i -= 1) {
            const rotor = self.rotors[i - 1];
            const offset = rotor.position;
            // Find the input that maps to current output
            for (0..26) |j| {
                if ((rotor.wiring[(j + offset) % 26] + 26 - offset) % 26 == current) {
                    current = @intCast(j);
                    break;
                }
            }
        }

        // Final plugboard
        return self.plugboard[current];
    }

    fn encrypt(self: *Enigma, input: []const u8, output: []u8) !void {
        for (input, 0..) |c, i| {
            if (c >= 'A' and c <= 'Z') {
                output[i] = @intCast('A' + self.encryptChar(c - 'A'));
            } else {
                output[i] = c; // Non-alphabetic characters unchanged
            }
        }
    }
};

pub fn main() !void {
    const allocator = std.heap.page_allocator;
    var enigma = try Enigma.init(allocator);
    enigma.setPositions("AAA"); // Starting position

    const stdout = std.io.getStdOut().writer();
    const input = "HELLO";
    var output: [5]u8 = undefined;

    try enigma.encrypt(input, &output);
    try stdout.print("Input: {s}\n", .{input});
    try stdout.print("Encrypted: {s}\n", .{output});
}
