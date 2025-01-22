// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

library Set {
    struct AddressSet {
        address[] values;
        mapping(address => uint256) positions;
    }

    error AddressSet_ValueNotFound(address value);
    error AddressSet_DuplicatedValue(address value);

    function insert(AddressSet storage _set, address _value) internal {
        if (_set.positions[_value] != 0)
            revert AddressSet_DuplicatedValue(_value);
        _set.values.push(_value);
        _set.positions[_value] = _set.values.length;
    }

    function erase(AddressSet storage _set, address _value) internal {
        uint256 p = _set.positions[_value];
        if (p == 0) revert AddressSet_ValueNotFound(_value);
        unchecked {
            _set.values[p - 1] = _set.values[_set.values.length - 1];
            _set.positions[_set.values[p - 1]] = p;
        }
        _set.values.pop();
        _set.positions[_value] = 0;
    }

    function hasValue(
        AddressSet storage _set,
        address _value
    ) internal view returns (bool) {
        return _set.positions[_value] != 0;
    }

    function isEmpty(AddressSet storage _set) internal view returns (bool) {
        return _set.values.length == 0;
    }

    function size(AddressSet storage _set) internal view returns (uint256) {
        return _set.values.length;
    }

    struct Uint256Set {
        uint256[] values;
        mapping(uint256 => uint256) positions;
    }

    error Uint256Set_ValueNotFound(uint256 value);
    error Uint256Set_DuplicatedValue(uint256 value);

    function insert(Uint256Set storage _set, uint256 _value) internal {
        if (_set.positions[_value] != 0)
            revert Uint256Set_DuplicatedValue(_value);
        _set.values.push(_value);
        _set.positions[_value] = _set.values.length;
    }

    function erase(Uint256Set storage _set, uint256 _value) internal {
        uint256 p = _set.positions[_value];
        if (p == 0) revert Uint256Set_ValueNotFound(_value);
        unchecked {
            _set.values[p - 1] = _set.values[_set.values.length - 1];
            _set.positions[_set.values[p - 1]] = p;
        }
        _set.values.pop();
        _set.positions[_value] = 0;
    }

    function hasValue(
        Uint256Set storage _set,
        uint256 _value
    ) internal view returns (bool) {
        return _set.positions[_value] != 0;
    }

    function isEmpty(Uint256Set storage _set) internal view returns (bool) {
        return _set.values.length == 0;
    }

    function size(Uint256Set storage _set) internal view returns (uint256) {
        return _set.values.length;
    }

    struct Bytes32Set {
        bytes32[] values;
        mapping(bytes32 => uint256) positions;
    }

    error Bytes32Set_ValueNotFound(bytes32 value);
    error Bytes32Set_OutBound(uint256 position);
    error Bytes32Set_DuplicatedValue(bytes32 value);

    function insert(Bytes32Set storage _set, bytes32 _value) internal {
        if (_set.positions[_value] != 0)
            revert Bytes32Set_DuplicatedValue(_value);
        _set.values.push(_value);
        _set.positions[_value] = _set.values.length;
    }

    function hasValue(
        Bytes32Set storage _set,
        bytes32 _value
    ) internal view returns (bool) {
        return _set.positions[_value] != 0;
    }

    function isEmpty(Bytes32Set storage _set) internal view returns (bool) {
        return _set.values.length == 0;
    }

    function size(Bytes32Set storage _set) internal view returns (uint256) {
        return _set.values.length;
    }
}
