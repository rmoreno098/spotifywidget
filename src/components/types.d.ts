export interface UserProfile {
    country: string;
    display_name: string;
    email: string;
    explicit_content: {
      filter_enabled: boolean;
      filter_locked: boolean;
    };
    external_urls: { spotify: string };
    followers: { href: string; total: number };
    href: string;
    id: string;
    images: Image[];
    product: string;
    type: string;
    uri: string;
  }
  
export interface Image {
    url: string;
    height: number;
    width: number;
  }
  
export interface Playlist {
    href: string;
    items: Item[];
    limit: number;
    next: string;
    offset: number;
    previous: string;
    total: number;
  }
  
export interface Item {
    collaborative: boolean;
    description: string;
    external_urls: { spotify: string };
    href: string;
    id: string;
    images: Image[];
    name: string;
    owner: Owner;
    primary_color: string;
    public: boolean;
    snapshot_id: string;
    tracks: Tracks;
    type: string;
    uri: string;
  }
  
export interface Tracks {
    href: string;
    total: number;
  }
  
export interface Owner {
    display_name: string;
    external_urls: { spotify: string };
    href: string;
    id: string;
    type: string;
    uri: string;
  }