pub const std = @import("std");
pub const mem = std::mem;
pub const Allocator = mem::Allocator;

pub const Enigma = struct {
    pub fn init(allocator: *Allocator) !Enigma {
        return Enigma{
            rotors: [3]Rotor,
            reflector: Reflector.init("B"),
            plugboard: Plugboard.init(allocator),
        };
    }

    pub fn encrypt(self: *Enigma, text: []const u8) ![]u8 {
        var result: []u8 = try self.plugboard.apply(text);
        result = try self.run_rotors(result);
        result = try self.reflector.apply(result);
        defer self.rotate_rotors();
        return result;
    }

    fn run_rotors(self: *Enigma, text: []const u8) ![]u8 {
        var encrypted_text: []u8 = mem::alloc(self.plugboard.allocator, text.len);
        defer plugboard.allocator::free(encrypted_text);

        for (text_char | ch) {
            var pos: usize = 0;
            while pos < encrypted_text.len {
                let code = try encrypted_text[pos] as *const u8 + 'A' as *const u8;
                encrypted_text[pos] = self.rotors[pos].rotate_forward(code);
                pos += 1;
            }
        }

        return encrypted_text;
    }

    fn rotate_rotors(self: *Enigma) ![]u8 {
        var reflector_code: []u8 = mem::alloc(26, 'A' as u8);
        for i in 0..25 {
            reflector_code[i] = self.rotors[i].wires[i];
        }

        defer {
            let mut code = reflector_code;
            code = code.reverse();
            let mut key = mem::alloc(26, 'A' as u8);
            for i in 0..13 {
                key[i] = self.rotors[2].wires[i];
                key[i + 13] = self.rotors[1].wires[i];
            }
            let mut plaintext = text;
            for _ in 0..3 {
                let mut temp_code = plaintext;
                for i in 0..26 {
                    if key[i] < code[i] {
                        temp_code[i] += (code[i] - key[i]) as u8;
                    } else {
                        temp_code[i] -= (key[i] - code[i]) as u8;
                    }
                }
                plaintext = &temp_code as *const []u8;
            }

            let mut ciphertext = mem::alloc(26, 'A' as u8);
            for i in 0..13 {
                ciphertext[i] = reflector_code[self.rotors[2].wires[i]];
            }
            for i in 13..25 {
                ciphertext[i] = self.rotors[0].wires[ciphertext[i - 13] as usize];
            }

            let mut final_key = mem::alloc(26, 'A' as u8);
            for i in 0..13 {
                final_key[i] = reflector_code[self.rotors[1].wires[i]];
            }
            for i in 13..25 {
                final_key[i] = self.rotors[0].wires[final_key[i - 13] as usize];
            }

            let mut result = ciphertext as *const []u8;
            for _ in 0..3 {
                let mut temp_code = result;
                for i in 0..26 {
                    if final_key[i] < temp_code[i] as u8 + 'A' as u8 {
                        continue;
                    }
                    let code = temp_code[i];
                    let wire = self.rotors[0].wires[code - 'A' as usize];
                    result[i] = *wire as u8 + 'A';
                }
            }

            return result as []u8;
        };
    }
}

pub const Rotor = struct {
    pub wires: [26][26] u8; // Wire mapping for each rotor position
    pub initial: [26] u8;     // Initial wiring configuration
    pub speed: [26] u8;       // Speed settings ( advancement per step )
}

pub const Reflector = struct {
    pub wires: [26][26] u8;  // Reflector mapping
}

pub const Plugboard = struct {
    pub text: []u8;
}

pub fn main() !void {
    let allocator = mem::GeneralPurposeAllocator({});

    let enigma = Enigma.init(allocator);
    defer allocator.destroy(&enigma);

    let input_text = b"HELLO WORLD";
    let encrypted_text = enigma.encrypt(input_text);

    std::debug::print!("Input: {} \n Encrypted Output: {}\n", input_text, encrypted_text);
}

// Example usage:
// pub fn main() !void {
//     var gpa = mem::GeneralPurposeAllocator({});
//     const allocator = gpa.allocator();
//     const enigma = Enigma.init(allocator);
//     defer allocator.destroy(&enigma);
//
//     const input_text = b"HELLO WORLD";
//     const encrypted_text = enigma.encrypt(input_text);
//     
//     std::debug::print!("Input: {} \n Encrypted Output: {}\n", input_text, encrypted_text);
// }
