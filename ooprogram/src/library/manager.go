//音乐管理
package library

import "errors"

type MusicManager struct {
	musics []MusicEntry
}

//工厂方法
func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

//返回的是MusicEntry类型变量的指针
func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removedMusic := &m.musics[index]

	if index < len(m.musics)-1 { //最后一个之前，用index+1(包括)后的数据追加到index-1之后
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 { //仅有一个元素
		m.musics = make([]MusicEntry, 0)
	} else { //最后一个元素
		m.musics = m.musics[:index-1]
	}

	return removedMusic
}

func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	return nil
}
